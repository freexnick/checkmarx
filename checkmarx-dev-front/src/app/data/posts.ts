"use server";
import { API_URI } from "@conf/config";
import { validateToken } from "app/auth/signIn";
import { redirect } from "next/navigation";
import { cookies } from "next/headers";

type formEntries = Record<string, number | FormDataEntryValue>;

async function fetchPost(postId: string) {
    if (!validateToken()) {
        redirect("/");
    }

    const token = cookies().get("session")?.value;

    try {
        const res = await fetch(`${API_URI}/posts/${postId}`, {
            method: "GET",
            cache: "no-store",
            headers: {
                "Content-Type": "application/json",
                Cookie: `session=${token || ""}`,
                Authorization: `Bearer ${token || ""}`,
            },
        });

        if (!res.ok) {
            throw new Error("Failed to fetch data");
        }

        return res.json();
    } catch (e) {
        console.error(e);
    }
}

async function fetchPosts() {
    if (!validateToken()) {
        redirect("/");
    }
    const token = cookies().get("session")?.value;

    try {
        const res = await fetch(`${API_URI}/posts`, {
            method: "GET",
            cache: "no-store",
            headers: {
                "Content-Type": "application/json",
                Cookie: `session=${token || ""}`,
                Authorization: `Bearer ${token || ""}`,
            },

            credentials: "include",
        });

        if (!res.ok) {
            throw new Error("Failed to fetch posts");
        }

        return res.json();
    } catch (e) {
        console.error(e);
    }
}

async function createPost(data: formEntries) {
    if (!validateToken()) {
        redirect("/");
    }

    const token = cookies().get("session")?.value;

    try {
        const res = await fetch(`${API_URI}/posts`, {
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
            throw new Error("Failed to create the post");
        }

        return res.statusText;
    } catch (e) {
        console.error(e);
    }
}

async function updatePost(data: formEntries) {
    if (!validateToken()) {
        redirect("/");
    }

    const token = cookies().get("session")?.value;

    try {
        const res = await fetch(`${API_URI}/posts/${data.id}`, {
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
            throw new Error("Failed to update the post");
        }

        return res.statusText;
    } catch (e) {
        console.error(e);
    }
}

async function deletePost(id: number) {
    if (!validateToken()) {
        redirect("/");
    }

    const token = cookies().get("session")?.value;

    try {
        const res = await fetch(`${API_URI}/posts/${id}`, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
                Cookie: `session=${token || ""}`,
                Authorization: `Bearer ${token || ""}`,
            },
            credentials: "include",
        });

        if (!res.ok) {
            throw new Error("Failed to delete the post");
        }

        return res.statusText;
    } catch (e) {
        console.error(e);
    }
}

export { fetchPost, fetchPosts, createPost, updatePost, deletePost };
