using System;
using System.Collections.Generic;

namespace ViewFinder
{
    class Program
    {
        static void Main(string[] args)
        {
            WindowClient client = new WindowClient("http://localhost:8090");

            Console.WriteLine("The following windows are above the browser in z order:\n");
            foreach (WindowInfo window in OpenWindowGetter.GetWindowsAboveBrowser())
            {
                IntPtr handle = window.WindowHandle;
                string title = window.WindowText;
                long topX = window.Info.rcWindow.Left;
                long topY = window.Info.rcWindow.Top;
                long botX = window.Info.rcWindow.Right;
                long botY = window.Info.rcWindow.Bottom;

                Console.WriteLine("{0}:      {1}\nTopX: {2}\nTopY: {3}\nBotX: {4}\nBotY: {5}\n\n", handle, title, topX, topY, botX, botY);

                client.SendWindowInfo(window).Wait();

            }
        }
    }
}