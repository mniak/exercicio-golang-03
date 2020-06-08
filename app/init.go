package app

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mniak/superlink-sdk"
	"github.com/revel/revel"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

var DB *sql.DB
var SuperLink *superlink.Client

func InitDB() {

	databaseHost, ok := revel.Config.String("database.host")
	if !ok {
		revel.AppLog.Fatal("Could not load database.host from config")
		return
	}
	connstring := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s sslmode=disable", databaseHost, "docker", "docker", "docker")

	var err error
	DB, err = sql.Open("postgres", connstring)
	if err != nil {
		revel.AppLog.Fatal("DB Error", err)
	}
}

func InitSuperLink() {
	baseurl, ok := revel.Config.String("superlink.baseurl")
	if !ok {
		revel.AppLog.Fatal("Could not load superlink.baseurl config")
		return
	}
	clientid, ok := revel.Config.String("superlink.clientid")
	if !ok {
		revel.AppLog.Fatal("Could not load superlink.clientid from config")
		return
	}
	clientsecret, ok := revel.Config.String("superlink.clientsecret")
	if !ok {
		revel.AppLog.Fatal("Could not load superlink.clientsecret from config")
		return
	}
	SuperLink = superlink.NewClient(baseurl, clientid, clientsecret)
	if revel.DevMode {
		SuperLink.SetDebug(true)
	}
}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	revel.OnAppStart(InitDB)
	revel.OnAppStart(InitSuperLink)
	// revel.OnAppStart(FillCache)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
