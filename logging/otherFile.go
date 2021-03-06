package logging

import "errors"

func Otherfunc() {
	Log.Trace("Something very low level.")
	Log.Debug("Useful debugging information.")
	Log.Info("Something noteworthy happened!")
	Log.Warn("You should probably take a look at this.")
	Log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	Log.Fatal("Bye. ", errors.New("wup"))
	// Calls panic() after logging
	//Log.Panic("I'm bailing.")
}
