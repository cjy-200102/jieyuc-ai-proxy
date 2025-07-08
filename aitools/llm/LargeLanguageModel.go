package llm

type LargeLanguageModel interface {
	Exec(string) string
}
