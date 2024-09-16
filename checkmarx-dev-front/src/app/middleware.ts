import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

const protectedRoutes = ["/*"];

export default async function middleware(req: NextRequest) {
    const path = req.nextUrl.pathname;
    const isProtectedRoute = protectedRoutes.includes(path);

    const cookie = cookies().get("session")?.value;
    const session = cookie?.length === 26;
    if (isProtectedRoute && !session) {
        return NextResponse.redirect(new URL("/", req.nextUrl));
    }

    return NextResponse.next();
}

export const config = {
    matcher: ["/((?!api|_next/static|_next/image|.*\\.png$).*)", "/api/:path*"],
};
