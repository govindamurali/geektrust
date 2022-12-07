package reader

type FilePathFetcher interface {
	GetFilePath() string
}

func GetConsoleFilePathFetcher() FilePathFetcher {
	return &consoleReader{}
}
