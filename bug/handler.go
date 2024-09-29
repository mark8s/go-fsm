package bug

import "fmt"

var (
	handlerBugDefault = Handler(func(args *Args) error {
		return nil
	})

	handlerBugNew = Handler(func(args *Args) error {
		return nil
	})

	handlerBugInProgress = Handler(func(args *Args) error {
		return nil
	})

	handlerBugResolved = Handler(func(args *Args) error {

		fmt.Println("状态变更为已解决")

		return nil
	})

	handlerBugReopened = Handler(func(args *Args) error {
		return nil
	})

	handlerBugRejected = Handler(func(args *Args) error {
		return nil
	})

	handlerBugClosed = Handler(func(args *Args) error {
		return nil
	})

	handlerBugSuspended = Handler(func(args *Args) error {
		return nil
	})
)
