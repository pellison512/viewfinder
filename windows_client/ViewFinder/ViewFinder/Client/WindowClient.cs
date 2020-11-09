using System;
using System.Collections.Generic;
using System.Net;
using System.Net.Http;
using System.Runtime.CompilerServices;
using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Threading.Tasks;
using System.Web;

public class WindowClient
{
	private class WindowRequest
	{
        public string windowText { get; }
        public int left { get; }
        public int top { get; }
        public  int right { get; }
        public int bottom { get; }

        public WindowRequest(string windowText, int left, int top, int right, int bottom)
        {
            this.windowText = windowText;
            this.left = left;
            this.top = top;
            this.right = right;
            this.bottom = bottom;
        }
    }

	static readonly HttpClient client = new HttpClient();
	static string baseURL;
	public WindowClient(string Url)
	{
		baseURL = Url;
	}
	public async Task SendWindowInfo(WindowInfo window)
	{
        List<WindowRequest> reqs = new List<WindowRequest>();
        
        WindowRequest req = ConvertWindowToReq(window);
        reqs.Add(req);
        Console.WriteLine("request: left: {0} right: {1}", req.left, req.right);

        string json = JsonSerializer.Serialize(reqs);
        Console.WriteLine("JSON IS: {0}", json);

        HttpContent content = new StringContent(json, Encoding.UTF8, "application/json");
		HttpResponseMessage response = await client.PostAsync(baseURL+"/windows", content);

	}
    public async Task SendWindowInfo(WindowInfo[] windows)
    {
        List<WindowRequest> reqs = new List<WindowRequest>();
        foreach(WindowInfo window in windows)
        {
            WindowRequest req = ConvertWindowToReq(window);
            Console.WriteLine("request: left: {0} right: {1}", req.left, req.right);
            reqs.Add(req);
        }
        
        string json = JsonSerializer.Serialize(reqs);

        HttpContent content = new StringContent(json, Encoding.UTF8, "application/json");
        HttpResponseMessage response = await client.PostAsync(baseURL+"/windows", content);
    }

    private WindowRequest ConvertWindowToReq(WindowInfo window)
    {
        string text = makeWindowNameFriendly(window.CleanBrowserText, window.IsBrowserWindow());
        
        return new WindowRequest(text, window.Info.rcWindow.Left, window.Info.rcWindow.Top, window.Info.rcWindow.Right, window.Info.rcWindow.Bottom);
    }
    //Removes non-browser names and replaces them with an id
    private string makeWindowNameFriendly(string windowText, bool isBrowserWindowName)
    {
        //TODO remove this, it's probably unecessary
        // return Uri.EscapeDataString(windowText);
        return windowText;

        //TODO implement this
       // return windowText;
    }

} 

