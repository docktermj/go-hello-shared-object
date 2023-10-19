package encrypt

type Encrypt interface {
	InitEncryption() error
	CloseEncryption() error
	Encrypt(rawText string, maxLen int) (string, error)
	EncryptDeterministic(rawText string, maxLen int) (string, error)
	Decrypt(rawText string, maxLen int) (string, error)
	DecryptDeterministic(rawText string, maxLen int) (string, error)
}
