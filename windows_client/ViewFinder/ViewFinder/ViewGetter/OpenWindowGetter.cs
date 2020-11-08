using System;
using System.Collections;
using System.Collections.Generic;
using System.Runtime.InteropServices;
using System.Text;
using System.Windows;
using HWND = System.IntPtr;

[StructLayout(LayoutKind.Sequential)]
public struct RECT
{
	public int Left, Top, Right, Bottom;

	public RECT(int left, int top, int right, int bottom)
	{
		Left = left;
		Top = top;
		Right = right;
		Bottom = bottom;
	}

	public RECT(System.Drawing.Rectangle r) : this(r.Left, r.Top, r.Right, r.Bottom) { }

	public int X
	{
		get { return Left; }
		set { Right -= (Left - value); Left = value; }
	}

	public int Y
	{
		get { return Top; }
		set { Bottom -= (Top - value); Top = value; }
	}

	public int Height
	{
		get { return Bottom - Top; }
		set { Bottom = value + Top; }
	}

	public int Width
	{
		get { return Right - Left; }
		set { Right = value + Left; }
	}

	public System.Drawing.Point Location
	{
		get { return new System.Drawing.Point(Left, Top); }
		set { X = value.X; Y = value.Y; }
	}

	public System.Drawing.Size Size
	{
		get { return new System.Drawing.Size(Width, Height); }
		set { Width = value.Width; Height = value.Height; }
	}

	public static implicit operator System.Drawing.Rectangle(RECT r)
	{
		return new System.Drawing.Rectangle(r.Left, r.Top, r.Width, r.Height);
	}

	public static implicit operator RECT(System.Drawing.Rectangle r)
	{
		return new RECT(r);
	}

	public static bool operator ==(RECT r1, RECT r2)
	{
		return r1.Equals(r2);
	}

	public static bool operator !=(RECT r1, RECT r2)
	{
		return !r1.Equals(r2);
	}

	public bool Equals(RECT r)
	{
		return r.Left == Left && r.Top == Top && r.Right == Right && r.Bottom == Bottom;
	}

	public override bool Equals(object obj)
	{
		if (obj is RECT)
			return Equals((RECT)obj);
		else if (obj is System.Drawing.Rectangle)
			return Equals(new RECT((System.Drawing.Rectangle)obj));
		return false;
	}

	public override int GetHashCode()
	{
		return ((System.Drawing.Rectangle)this).GetHashCode();
	}

	public override string ToString()
	{
		return string.Format(System.Globalization.CultureInfo.CurrentCulture, "{{Left={0},Top={1},Right={2},Bottom={3}}}", Left, Top, Right, Bottom);
	}
}


[StructLayout(LayoutKind.Sequential)]
public struct WINDOWINFO
{
	public uint cbSize;
	public RECT rcWindow;
	public RECT rcClient;
	public uint dwStyle;
	public uint dwExStyle;
	public uint dwWindowStatus;
	public uint cxWindowBorders;
	public uint cyWindowBorders;
	public ushort atomWindowType;
	public ushort wCreatorVersion;

	public WINDOWINFO(Boolean? filler) : this()   // Allows automatic initialization of "cbSize" with "new WINDOWINFO(null/true/false)".
	{
		cbSize = (UInt32)(Marshal.SizeOf(typeof(WINDOWINFO)));
	}
}

public class WindowInfo
{
	public string WindowText;
	public WINDOWINFO Info;
	public HWND WindowHandle;
}

	public static class OpenWindowGetter
{
	private delegate bool EnumWindowsProc(HWND hWnd, int lParam);

	/*
	public Class1()
	{
	}*/
	public static IDictionary<HWND, WindowInfo> GetWindowsAboveBrowser()
	{
		HWND shellWindow = GetShellWindow();
		//Dictionary<HWND, string> windows = new Dictionary<HWND, string>();

		int lastChromeZCount = -1;
		int counter = 0;

		List<WindowInfo> allVisibleWindows = new List<WindowInfo>();

		//Go through all the windows visible and add them to a list
		EnumWindows(delegate (HWND hWnd, int lParam)
		{
			if (hWnd == shellWindow)
			{
				return true;
			}
			if (!IsWindowVisible(hWnd))
			{
				return true;
			}

			int length = GetWindowTextLength(hWnd);
			if (length == 0)
			{
				return true;
			}

			

			WINDOWINFO info = new WINDOWINFO();
			info.cbSize = (uint)Marshal.SizeOf(info);
			GetWindowInfo(hWnd, ref info);

			StringBuilder builder = new StringBuilder(length);
			GetWindowText(hWnd, builder, length + 1);

			WindowInfo winInfo = new WindowInfo();
			winInfo.Info = info;
			winInfo.WindowText = builder.ToString();
			winInfo.WindowHandle = hWnd;
			allVisibleWindows.Add(winInfo);

            if(builder.ToString().ToLower().Contains("chrome"))
            {
				lastChromeZCount = counter;
            }
			counter++;

			return true;
		}, 0);

		//go through the visible windows and trim them to only those above the lowest chrome window
		Dictionary<HWND, WindowInfo> windowsAboveBrowser = new Dictionary<HWND, WindowInfo>();
		for (int i=0; i< allVisibleWindows.Count; i++) {
			WindowInfo window = allVisibleWindows[i];
			
			if(i < lastChromeZCount)
            {
				windowsAboveBrowser[window.WindowHandle] = window;
            }
        }

		return windowsAboveBrowser;
	}

	[DllImport("USER32.DLL")]
	private static extern bool GetWindowInfo(IntPtr hwnd, ref WINDOWINFO pwi);

	[DllImport("USER32.DLL")]
	private static extern bool IsWindowVisible(HWND hWnd);

	[DllImport("USER32.DLL")]
	private static extern IntPtr GetShellWindow();

	[DllImport("USER32.DLL")]
	private static extern bool EnumWindows(EnumWindowsProc enumFunc, int lParam);

	[DllImport("USER32.DLL")]
	private static extern int GetWindowTextLength(HWND hWnd);

	[DllImport("USER32.DLL")]
	private static extern int GetWindowText(HWND hWnd, StringBuilder lpString, int nMaxCount);
}

