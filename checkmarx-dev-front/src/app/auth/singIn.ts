"use server";
import { API_URI } from "@conf/config";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

type formEntries = Record<string, number | FormDataEntryValue | string>;

async function signUser(data: formEntries) {
    try {
        const res = await fetch(`${API_URI}/auth/signin`, {
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
    } catch (e) {
        console.error(e);
    }
}

async function fetchUser() {
    if (!validateToken()) {
        redirect("/");
    }

    const token = cookies().get("session")?.value;

    try {
        const res = await fetch(`${API_URI}/auth`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                Cookie: `session=${token || ""}`,
                Authorization: `Bearer ${token || ""}`,
            },

            credentials: "include",
        });

        const user = await res.json();

        if (!res.ok) {
            if (res.status === 404) {
                throw new Error(`${res.status}:${res.statusText}`);
            }
        }
        return user;
    } catch (e) {
        console.error(e);
    }
}

function validateToken() {
    const cookie = cookies().get("session")?.value;
    return cookie && cookie.length >= 26;
}

export { signUser, validateToken, fetchUser };
