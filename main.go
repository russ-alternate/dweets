package main

// -------------------------------------------------------------------
// get the last dweet for a thing
// -------------------------------------------------------------------
import (
	"fmt"
)

var dweetOp = "get_last"

//var dweetOp = "get_all"

var dweetThing = "bent-train"

//var dweetThing = "bent-bum"

func main() {

	if dweetOp == "get_last" {
		fmt.Println("in section get_last")

		api := Dweetio{}
		res, err := api.GetLastDweetFor(dweetThing)

		if err != nil {
			//	fmt.Println(err.Error())
			fmt.Println("Error in GetLastDweetFor")
			fmt.Println(err.Error())
		}

		if res.This == "failed" {
			fmt.Println("GetLastDweetFor - call to dweet.io failed")
		} else {
			fmt.Println(res)
		}

		//				api.ReturnError(err)
		//		res, err := api.GetLastDweetFor(dweetThing)

	} else if dweetOp == "get_last_locked" {
		fmt.Println("in section get_last_locked")

		/*		api := dweetio.Dweetio{Key: "myPrivateLockKey"}
					res, err := api.GetLastDweetFor(dweetThing)
					if err != nil {
						//			t.Error("Error in the GetLastDweetFor")
						fmt.Println(err.Error())
					}
					fmt.Println(res)

				} else if dweetOp == "get_all" {
					fmt.Println("in section get_all")

					api := dweetio.Dweetio{}
					res, err := api.GetDweetsFor(dweetThing)
					if err != nil {
						//			t.Error("Error in the GetDweetsFor")
						fmt.Println(err.Error())
					}
					fmt.Println(res)

				} else if dweetOp == "get_all_locked" {
					fmt.Println("in section get_all_locked")

					api := dweetio.Dweetio{Key: "myPrivateLockKey"}
					res, err := api.GetDweetsFor(dweetThing)
					if err != nil {
						//			t.Error("Error in the GetDweetsFor")
						fmt.Println(err.Error())
					}
					fmt.Println(res)
		*/
	} else {
		fmt.Println("nuthin happening")
	}

	//	if err != nil {
	//		//		t.Error("Error in the GetLastDweetFor")
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println(res)

}

//api := dweetio.Dweetio{}
//res, err := api.GetLastDweetFor(dweet_thing)
//
//if err != nil {
//	//		t.Error("Error in the GetLastDweetFor")
//	fmt.Println(err.Error())
//}
//fmt.Println(res)
//

/*
// -------------------------------------------------------------------
// get the last dweet for a LOCKED thing
// -------------------------------------------------------------------
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{Key: "myPrivateLockKey"}
	res, err := api.GetLastDweetFor("godweetio") //Change this value with the name of your thing

	if err != nil {
		t.Error("Error in the GetLastDweetFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}

// -------------------------------------------------------------------
// get all the dweets for a thing - limited to 500
// -------------------------------------------------------------------
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{}
	res, err := api.GetDweetsFor("godweetio") //Change this value with the name of your thing

	if err != nil {
		t.Error("Error in the GetDweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}


// -------------------------------------------------------------------
// get all the dweets for a locked thing - limited to 500
// -------------------------------------------------------------------
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{Key: "myPrivateLockKey"}
	res, err := api.GetDweetsFor("godweetio") //Change this value with the name of your thing

	if err != nil {
		t.Error("Error in the GetDweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}


// -------------------------------------------------------------------
// post new data for a thing
// -------------------------------------------------------------------
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{}
	res, err := api.DweetFor("godweetio", "{ \"lorem\": \"ipsum\" }")

	if err != nil {
		t.Error("Error in the DweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}


// -------------------------------------------------------------------
// post new data for a locked thing
// -------------------------------------------------------------------
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{Key: "myPrivateLockKey"}
	res, err := api.DweetFor("godweetio", "{ \"lorem\": \"ipsum\" }")

	if err != nil {
		t.Error("Error in the DweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}


// -------------------------------------------------------------------
// set an alert for a locked thing
//
// first parmeter - thing (name of your thing)
// second parmeter - recipients (a list of strings)
// third parmeter - condition - a javascript expression to evaluate
//								the data in a dweet and to return
//								whether or not an alert should be sent
// -------------------------------------------------------------------
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{Key: "myLockKey"}
	res, err := api.SetAlertFor("godweetio", []string{"test@godweetio.com"}, "if(dweet.alertValue > 10) return 'TEST: Greater than 10'; if(dweet.alertValue < 10) return 'TEST: Less than 10';")

	if err != nil {
		t.Error("Error in the DweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}


// -------------------------------------------------------------------
// get an alert for a locked thing
// -------------------------------------------------------------------
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{Key: "myLockKey"}
	res, err := api.GetAlertFor("godweetio")

	if err != nil {
		t.Error("Error in the GetAlertFor")
		fmt.Println(err)
	}

	fmt.Println(res)
}


// -------------------------------------------------------------------
// remove an alert for a locked thing
// -------------------------------------------------------------------
import (
	"fmt"
	"github.com/vitorleal/go-dweetio"
)

func main() {
	api := Dweetio{Key: "myLockKey"}
	res, err := api.RemoveAlertFor("godweetio")

	if err != nil {
		t.Error("Error in the DweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
*/
