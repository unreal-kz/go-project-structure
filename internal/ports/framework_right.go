package ports

type DbProt interface {
	CloseDbConnection()
	AddToHistory(answer int32, operation string) error
}
