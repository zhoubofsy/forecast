package speaker

type say_if interface{
	Init(ak string, sec string, cuid string) error
	Create_Token() (string, error)
	Do_Text2audio(text string) (error)
}

