
<script setup>
import { ref, computed, watch } from "vue";

/* ----------------------------------
   PROPS
---------------------------------- */
const props = defineProps({
  modelValue: { type: Boolean, default: false },
  playlists: { type: Array, default: () => [] },
  alarm: { type: Object, default: null },  // nullable
});
const emit = defineEmits(["update:modelValue", "save"]);

/* ----------------------------------
   DIALOG CONTROL
---------------------------------- */
const isOpen = computed({
  get: () => props.modelValue,
  set: (v) => emit("update:modelValue", v),
});

/* ----------------------------------
   ENUM VALUES FROM BACKEND
---------------------------------- */
const weekdays = [
  { short: "M",  value: "WEEKDAY_MONDAY", id: 1 },
  { short: "T",  value: "WEEKDAY_TUESDAY", id: 2 },
  { short: "W",  value: "WEEKDAY_WEDNESDAY", id: 3 },
  { short: "T",  value: "WEEKDAY_THURSDAY", id: 4 },
  { short: "F",  value: "WEEKDAY_FRIDAY", id: 5 },
  { short: "S",  value: "WEEKDAY_SATURDAY", id: 6 },
  { short: "S",  value: "WEEKDAY_SUNDAY", id: 7 },
];

/* ----------------------------------
   CREATE A VALID PROTO ALARM STRUCT
---------------------------------- */
function emptyAlarm() {
  return {
    id: "",                             // backend assigns ID on create
    label: "",
    time: "07:00",
    repeatDays: [],
    enabled: true,
    warmupDuration: "0s",               // protobuf duration
    ledTarget: { r: 255, g: 200, b: 150 },
    playableId: "",
  };
}

/* ----------------------------------
   EDIT STATE
---------------------------------- */
const editable = ref(emptyAlarm());
const editableColor = ref({ r: 255, g: 200, b: 150 });

const isEditMode = computed(() => !!editable.value.id);

/* warmup slider stored in minutes,
   but backend wants protobuf duration "Xs"
*/
const warmupMinutes = ref(0);

/* ----------------------------------
   INIT WHEN OPENING DIALOG
---------------------------------- */
function init() {
  const src = props.alarm ?? emptyAlarm();

  editable.value = JSON.parse(JSON.stringify(src));

  // parse warmup "Xs" → minutes
  if (typeof src.warmupDuration === "string") {
    const seconds = parseInt(src.warmupDuration.replace("s", "")) || 0;
    warmupMinutes.value = Math.round(seconds / 60);
  }

  editableColor.value = {
    r: src.ledTarget?.r ?? 255,
    g: src.ledTarget?.g ?? 200,
    b: src.ledTarget?.b ?? 150,
  };
}

watch(() => props.alarm, init, { immediate: true });
watch(isOpen, (v) => v && init());

/* Keep LED color linked */
watch(editableColor, (rgb) => {
  editable.value.ledTarget = { ...rgb };
});

/* ----------------------------------
   CLOSE & SAVE
---------------------------------- */
function close() {
  isOpen.value = false;
}

function save() {
  const out = JSON.parse(JSON.stringify(editable.value));

  // Convert minutes → "Xs" protobuf duratio
  out.warmupDuration = {
  seconds: BigInt(warmupMinutes.value * 60),
  nanos: 0,
};

console.log("saving", out)
  emit("save", out);
  close();
}
</script>

<template>
  <v-dialog v-model="isOpen" max-width="500">
    <v-card class="editor-card">
      <v-card-title class="text-h6">
        {{ isEditMode ? 'Edit Alarm' : 'Add Alarm' }}
      </v-card-title>

      <v-card-text class="d-flex flex-column gap-4">

        <!-- NAME -->
        <v-text-field
          v-model="editable.label"
          label="Name"
          variant="outlined"
          density="comfortable"
        />

        <!-- TIME -->
        <v-text-field
          v-model="editable.time"
          label="Time"
          type="time"
          variant="outlined"
          density="comfortable"
        />

        <!-- REPEAT DAYS -->
        <div>
          <div class="text-body-2 mb-1">Repeat Days</div>

          <v-btn-toggle
            v-model="editable.repeatDays"
            multiple
            variant="default"
            color="var(--color-primary)"
            class="day-toggle"
          >
            <v-btn
              v-for="d in weekdays"
              :key="d.value"
              :value="d.id"
              size="small"
              class="day-btn"
            >
              {{ d.short }}
            </v-btn>
          </v-btn-toggle>
        </div>

        <!-- WARMUP (in minutes) -->
        <div>
          <div class="text-body-2 mb-1">
            Warmup Duration: {{ warmupMinutes }} min
          </div>
          <v-slider
            v-model="warmupMinutes"
            :min="0"
            :max="60"
            :step="5"
            color="var(--color-primary)"
          />
        </div>

        <!-- LED COLOR PICKER -->
        <div class="mb-8">
          <div class="text-body-2 mb-1">LED Color</div>

          <v-color-picker
            v-model="editableColor"
            mode="rgb"
            hide-mode-switch
            canvas-height="120"
          />
        </div>

        <!-- PLAYLIST -->
        <v-select
          v-model="editable.playableId"
          label="Playlist"
          :items="playlists"
          item-title="label"
          item-value="id"
          variant="outlined"
          density="comfortable"
        />
      </v-card-text>

      <v-card-actions class="d-flex justify-end">
        <v-btn variant="text" @click="close">Cancel</v-btn>
        <v-btn color="var(--color-primary)" @click="save">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<style scoped>
.editor-card {
  background-color: var(--color-bg);
  color: var(--color-surface);
  border-radius: 12px;
}
.day-toggle {
  display: flex;
  justify-content: space-between;
  width: 100%;
}
.day-btn {
  background-color: var(--inactive-bg);
}
.v-color-picker {
  border: 1px solid var(--color-border);
  border-radius: 8px;
}
</style>
