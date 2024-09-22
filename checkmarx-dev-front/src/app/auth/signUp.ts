"use server";
import { API_URI } from "@conf/config";
import { cookies } from "next/headers";

export async function signUpUser(data: Record<string, FormDataEntryValue>) {
    "use server";
    try {
        const res = await fetch(`${API_URI}/auth/signup`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),

            credentials: "include",
        });

        const user = await res.json();
        if (!res.ok) {
            return { status: res.status, message: user };
        }

        if (user?.session?.token) {
            cookies().set("session", user.session.token, {
                expires: new Date(user.session.expires_at),
                path: "/",
            });
        }

        return user;
    } catch (err) {
        console.error(err);
    }
}
