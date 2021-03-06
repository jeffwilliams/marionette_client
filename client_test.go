package marionette_client

import (
	"testing"
	"time"
)

const (
	TARGET_URL        = "http://www.abola.pt/"
	ID_SELECTOR       = "clubes-hp"
	CSS_SELECTOR_LI   = "li"
	ID_SELECTOR_INPUT = "topo_txtPesquisa"
	TIMEOUT           = 10000 // milliseconds
)

var client *Client

func init() {
	client = NewClient()
	client.Transport(&MarionetteTransport{})
	RunningInDebugMode = true
}

// we don't want parallel execution we need sequence.
func TestInit(t *testing.T) {
	t.Run("sequence", func(t *testing.T) {
		t.Run("NewSessionTest", NewSessionTest)
		t.Run("GetSessionIDTest", GetSessionIDTest)
		t.Run("GetPageTest", GetPageTest)
		t.Run("CurrentUrlTest", CurrentUrlTest)

		t.Run("GetCookiesTest", GetCookiesTest)
		t.Run("GetCookieTest", GetCookieTest)

		t.Run("GetSessionCapabilitiesTest", GetSessionCapabilitiesTest)
		t.Run("ScreenshotTest", ScreenshotTest)

		t.Run("LogTest", LogTest)
		t.Run("GetLogsTest", GetLogsTest)

		t.Run("SetContextTest", SetContextTest)
		t.Run("GetContextTest", GetContextTest)

		t.Run("GetPageSourceTest", GetPageSourceTest)

		t.Run("SetScriptTimoutTest", SetScriptTimoutTest)
		t.Run("SetPageTimoutTest", SetPageTimoutTest)
		t.Run("SetSearchTimoutTest", SetSearchTimoutTest)

		t.Run("PageSourceTest", PageSourceTest)

		t.Run("ExecuteScriptWithoutFunctionTest", ExecuteScriptWithoutFunctionTest)
		t.Run("ExecuteScriptTest", ExecuteScriptTest)
		t.Run("ExecuteScriptWithArgsTest", ExecuteScriptWithArgsTest)

		t.Run("GetTitleTest", GetTitleTest)

		t.Run("FindElementTest", FindElementTest)

		t.Run("SendKeysTest", SendKeysTest)
		t.Run("FindElementsTest", FindElementsTest)

		t.Run("CurrentChromeWindowHandleTest", CurrentChromeWindowHandleTest)
		t.Run("WindowHandlesTest", WindowHandlesTest)

		t.Run("NavigatorMethodsTest", NavigatorMethodsTest)

		t.Run("WaitForUntilIntegrationTest", WaitForUntilIntegrationTest)

		t.Run("PromptTest", PromptTest)
		t.Run("AlertTest", AlertTest)

		// test expected.go
		t.Run("NotPresentTest", NotPresentTest)

		t.Run("WindowSizeTest", WindowSizeTest)

		t.Run("DeleteSessionTest", DeleteSessionTest)

		// test QuitApplication
		t.Run("NewSessionTest", NewSessionTest)
		t.Run("QuitTest", QuitTest)
	})
}

/*********/
/* tests */
/*********/

func NewSessionTest(t *testing.T) {
	err := client.Connect("", 0)
	if err != nil {
		t.Fatalf("%#v", err)
	}
	t.Log("got here")
	r, err := client.NewSession("", nil)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func GetSessionIDTest(t *testing.T) {
	if client.SessionId != client.SessionID() {
		t.Fatalf("SessionId differ...")
	}

	t.Log("session is : ", client.SessionId)
}

func GetPageTest(t *testing.T) {
	r, err := client.Navigate(TARGET_URL)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func CurrentUrlTest(t *testing.T) {
	url, err := client.CurrentUrl()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	if url != TARGET_URL {
		t.Fatalf("Current Url %v not equal to target url %v", url, TARGET_URL)
	}

}

func GetCookiesTest(t *testing.T) {
	r, err := client.Cookies()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func GetCookieTest(t *testing.T) {
	r, err := client.Cookie("abolaCookie")
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

//func TestConnectWithActiveConnection(t *testing.T) {
//	err := client.Connect("", 0)
//	if err == nil {
//		t.Fatalf("%#v", err)
//	}
//
//	t.Log("No Error..")
//}

func GetSessionCapabilitiesTest(t *testing.T) {
	r, err := client.Capabilities()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.BrowserName)
}

func ScreenshotTest(t *testing.T) {
	_, err := client.Screenshot()
	if err != nil {
		t.Fatal(err)
	}

	//this print ise a problem for travis builds, since it can surpass the 4 MB of log size.
	// don't print the base64 encoded image.
	//println(base64encoded)
}

// working
func LogTest(t *testing.T) {
	r, err := client.Log("message testing", "warning")
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func GetLogsTest(t *testing.T) {
	r, err := client.Logs()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func SetContextTest(t *testing.T) {
	r, err := client.SetContext(Context(CHROME))
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)

	r, err = client.SetContext(Context(CONTENT))
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func GetContextTest(t *testing.T) {
	r, err := client.Context()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func GetPageSourceTest(t *testing.T) {
	r, err := client.SetContext(Context(CHROME))
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)

	r, err = client.SetContext(Context(CONTENT))
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func SetScriptTimoutTest(t *testing.T) {
	r, err := client.SetScriptTimeout(TIMEOUT)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func SetPageTimoutTest(t *testing.T) {
	r, err := client.SetPageTimeout(TIMEOUT)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func SetSearchTimoutTest(t *testing.T) {
	r, err := client.SetSearchTimeout(TIMEOUT)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func PageSourceTest(t *testing.T) {
	_, err := client.PageSource()
	if err != nil {
		t.Fatalf("%#v", err)
	}
}

func ExecuteScriptWithoutFunctionTest(t *testing.T) {
	script := "return (document.readyState == 'complete');"
	args := []interface{}{}
	r, err := client.ExecuteScript(script, args, TIMEOUT, false)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func ExecuteScriptTest(t *testing.T) {
	script := "function testMyGoMarionetteClient() { return 'yes'; } return testMyGoMarionetteClient();"
	args := []interface{}{}
	r, err := client.ExecuteScript(script, args, TIMEOUT, false)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func ExecuteScriptWithArgsTest(t *testing.T) {
	script := "function testMyGoMarionetteClientArgs(a, b) { return a + b; }; return testMyGoMarionetteClientArgs(arguments[0], arguments[1]);"
	args := []interface{}{1, 3}
	r, err := client.ExecuteScript(script, args, TIMEOUT, false)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func GetTitleTest(t *testing.T) {
	title, err := client.Title()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(title)

}
func FindElementTest(t *testing.T) {
	element, err := client.FindElement(By(ID), ID_SELECTOR)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(element.Id())
	t.Log(element.Enabled())
	t.Log(element.Selected())
	t.Log(element.Displayed())
	t.Log(element.TagName())
	t.Log(element.Text())
	t.Log(element.Attribute("id"))
	t.Log(element.CssValue("text-decoration"))
	rect, err := element.Rect()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(rect)

	// location
	point, err := element.Location()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Logf("x: %f, y: %f", point.X, point.Y)

	//size
	size, err := element.Size()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Logf("w: %f, h: %f", size.Width, size.Height)

	// screenshot of node element
	_, err = element.Screenshot()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	collection, err := element.FindElements(By(CSS_SELECTOR), CSS_SELECTOR_LI)
	if 18 != len(collection) {
		t.FailNow()
	}

	t.Logf("%T %#v", collection, collection)
}

func SendKeysTest(t *testing.T) {
	e, err := client.FindElement(By(ID), ID_SELECTOR_INPUT)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	e.SendKeys("teste")
}

func FindElementsTest(t *testing.T) {
	elements, err := client.FindElements(By(CSS_SELECTOR), CSS_SELECTOR_LI)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(len(elements))
}

func CurrentChromeWindowHandleTest(t *testing.T) {
	r, err := client.CurrentChromeWindowHandle()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func WindowHandlesTest(t *testing.T) {
	w, err := client.CurrentWindowHandle()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(w)

	r, err := client.WindowHandles()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	for _, w := range r {
		err := client.SwitchToWindow(w)
		if err != nil {
			t.Fatalf("%#v", err)
		}

		time.Sleep(time.Duration(time.Second))
	}

	// return to original window.
	client.SwitchToWindow(w)
}

func NavigatorMethodsTest(t *testing.T) {
	client.SetContext(Context(CONTENT))
	url1 := "https://www.google.pt/"
	url2 := "https://www.bing.com/"

	client.Navigate(url1)
	sleep := time.Duration(2) * time.Second
	time.Sleep(sleep)

	client.Navigate(url2)
	time.Sleep(sleep)

	client.Back()
	client.Refresh()
	time.Sleep(sleep)

	firstUrl, err := client.CurrentUrl()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	if firstUrl != url1 {
		t.Fatalf("Expected url %v - received url %v", url1, firstUrl[0:len(url1)])
	}

	client.Forward()
	secondUrl, err := client.CurrentUrl()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	if secondUrl != url2 {
		t.Fatalf("Expected url %v - received url %v", url2, secondUrl[:len(url2)])
	}
}

func PromptTest(t *testing.T) {
	client.Get("http://www.abola.pt")
	var text string = "marionette is cool or what - prompt?"
	var script string = "prompt('" + text + "');"
	args := []interface{}{}

	r, err := client.ExecuteScript(script, args, TIMEOUT, false)
	if err != nil {
		t.Fatalf("%#v", err)
	}
/* FIXME: changed to parameter text in firefox 55.0a1 branch
	err = client.SendKeysToDialog("yeah!")
	if err != nil {
		t.Fatalf("%#v", err)
	}
*/
	time.Sleep(time.Duration(5) * time.Second)

	err = client.DismissDialog()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func AlertTest(t *testing.T) {
	client.Get("http://www.abola.pt")
	var text string = "marionette is cool or what?"
	var script string = "alert('" + text + "');"
	args := []interface{}{}
	r, err := client.ExecuteScript(script, args, TIMEOUT, false)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	textFromdialog, err := client.TextFromDialog()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	if textFromdialog != text {
		t.Fatalf("Text in dialog differ. expected: %v, textfromdialog: %v", text, textFromdialog)
	}

	time.Sleep(time.Duration(5) * time.Second)

	err = client.AcceptDialog()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}

func WindowSizeTest(t *testing.T) {
	size, err := client.WindowSize()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Logf("w: %v, h: %v", size.Width, size.Height)

/* FIXME: SetWindowSize hangs on travis ci with XVFB start on before script. tests work localy
	newSize := &Size{Width: size.Width / 2, Height: size.Height / 2}
	rv, err := client.SetWindowSize(newSize)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Logf("new w: %v, new h: %v", rv.Width, rv.Height)

	err = client.MaximizeWindow()
	if err != nil {
		t.Fatalf("%#v", err)
	}
*/
}

// working - if called before other tests all hell will break loose
//func TestCloseWindow(t *testing.T) {
//	r, err := client.CloseWindow()
//	if err != nil {
//		t.Fatalf("%#v", err)
//	}
//
//	t.Log(r.Value)
//}

// working - if called before other tests all hell will break loose
func DeleteSessionTest(t *testing.T) {
	err := client.DeleteSession()
	if err != nil {
		t.Fatalf("%#v", err)
	}
}

func QuitTest(t *testing.T) {
	r, err := client.QuitApplication()
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Log(r.Value)
}
