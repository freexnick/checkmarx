"use server";
import { API_URI } from "@conf/config";
import { validateToken } from "app/auth/signIn";
import { redirect } from "next/navigation";
import { cookies } from "next/headers";

type formEntries = Record<string, number | FormDataEntryValue>;

async function createComment(data: formEntries) {
    if (!validateToken()) {
        redirect("/");
    }

    const token = cookies().get("session")?.value;

    try {
        const res = await fetch(`${API_URI}/comments`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Cookie: `session=${token || ""}`,
                Authorization: `Bearer ${token || ""}`,
            },
            body: JSON.stringify(data),
            credentials: "include",
        });

        if (!res.ok) {
            throw new Error("Failed to create comment");
        }

        return res.statusText;
    } catch (e) {
        console.error(e);
    }
}

async function updateComment(data: formEntries) {
    if (!validateToken()) {
        redirect("/");
    }

    const token = cookies().get("session")?.value;

    try {
        const res = await fetch(`${API_URI}/comments/${data.id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                Cookie: `session=${token || ""}`,
                Authorization: `Bearer ${token || ""}`,
            },
            body: JSON.stringify(data),
            credentials: "include",
        });

        if (!res.ok) {
            throw new Error("Failed to update comment");
        }

        return res.statusText;
    } catch (e) {
        console.error(e);
    }
}

export { createComment, updateComment };
