import "./globals.css";
import { ReactNode } from "react";
import Providers from "./context/Providers";

interface LayoutProps {
    children: ReactNode;
}

export default async function RootLayout({ children }: LayoutProps) {
    return (
        <html lang="en">
            <head>
                <title>Posts</title>
                <meta name="viewport" content="width=device-width, initial-scale=1" />
            </head>
            <body className="bg-gray-100 min-h-screen">
                <Providers>
                    <main>{children}</main>
                </Providers>
            </body>
        </html>
    );
}
