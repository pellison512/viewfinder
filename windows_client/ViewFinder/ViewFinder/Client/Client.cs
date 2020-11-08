using System;
using System.Net.Http;
using System.Runtime.CompilerServices;
using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Threading.Tasks;
public class Client
{
	private class WindowRequest
	{
        private readonly string windowText;
        private readonly int left;
        private readonly int top;
        private readonly int right;
        private readonly int bottom;

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
	public Client(string Url)
	{
		baseURL = Url;
	}
	public async Task SendWindowInfo(WindowInfo window)
	{
        WindowRequest req = ConvertWindowToReq(window);
        string json = JsonSerializer.Serialize(req);

        HttpContent content = new StringContent(json, Encoding.UTF8, "application/json");
		HttpResponseMessage resposne = await client.PostAsync(baseURL, content);
	}

	private WindowRequest ConvertWindowToReq(WindowInfo window)
    {
		return new WindowRequest(window.WindowText, window.Info.rcWindow.Left, window.Info.rcWindow.Top, window.Info.rcWindow.Right, window.Info.rcWindow.Bottom);
    }
}

