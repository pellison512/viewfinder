using System;
using System.Collections.Generic;

namespace ViewFinder
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("The following windows are above the browser in z order:\n");
            foreach (KeyValuePair<IntPtr, WindowInfo> window in OpenWindowGetter.GetWindowsAboveBrowser())
            {
                IntPtr handle = window.Key;
                string title = window.Value.WindowText;
                long topX = window.Value.Info.rcWindow.Left;
                long topY = window.Value.Info.rcWindow.Top;
                long botX = window.Value.Info.rcWindow.Right;
                long botY = window.Value.Info.rcWindow.Bottom;



                Console.WriteLine("{0}:      {1}\nTopX: {2}\nTopY: {3}\nBotX: {4}\nBotY: {5}\n\n", handle, title, topX, topY, botX, botY);
             

            }
        }
    }
}