package sources

import "fmt"

type LocalFileAdapter struct{}

func (l *LocalFileAdapter) GetPath(path string) (string, error) {
	// TODO: add check if file exists, yadda yadda
	return fmt.Sprintf("file://%v", path), nil
}

func (l *LocalFileAdapter) GetMultiPaths(paths ...string) ([]string, error) {
	ret := make([]string, 0, len(paths))
	for _, v := range paths {
		pp, err := l.GetPath(v)
		if err != nil {
			return nil, err
		}

		ret = append(ret, pp)
	}

	return ret, nil
}
