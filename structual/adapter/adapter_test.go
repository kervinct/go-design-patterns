package structual

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello world!"

	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}
	returnedMsg := adapter.PrintStored()
	if returnedMsg != "Legacy Printer: Adapter: Hello world!\n" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()
	if returnedMsg != "Hello world!" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
}
