export default defineNuxtPlugin((nuxtApp) => {
  console.log("[suppress-warnings] plugin loaded");

  nuxtApp.vueApp.config.warnHandler = (msg, instance, trace) => {
    // comment this out once you’re sure it triggers
    // console.log("[suppress-warnings] intercepted warn:", msg);
    // do nothing -> warning swallowed
  };
});
