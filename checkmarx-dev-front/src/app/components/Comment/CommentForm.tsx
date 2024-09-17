"use client";
import { useUser } from "app/context/User";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { createComment, updateComment } from "app/data/comments";
import { CommentState } from "@t/index";

interface CommentFormProps {
    postID: number;
    commentState: CommentState | null;
    onCommentSubmit?: (comment: { name: string; content: string }) => void;
    stopEditing: () => void;
}

export default function CommentForm({ postID, commentState, stopEditing }: CommentFormProps) {
    const { userRef } = useUser();
    const router = useRouter();
    const [content, setContent] = useState("");

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const formData = new FormData(event.currentTarget);
        const user_id = userRef?.current?.id as number;
        const data = Object.fromEntries(formData.entries()) as Record<string, FormDataEntryValue | number>;
        data.author_id = user_id;
        data.post_id = postID;

        let submission = createComment;

        if (commentState?.commentID) {
            data.id = commentState.commentID;
            submission = updateComment;
        }

        try {
            const res = await submission(data);

            if (res) {
                stopEditing();
                setContent("");
                router.refresh();
            }
        } catch (e) {
            console.error(e);
        }
    };

    useEffect(() => {
        setContent(commentState?.commentMessage || "");
    }, [commentState]);

    return (
        <form onSubmit={handleSubmit} className="space-y-4 mt-6">
            <div className="space-y-2">
                <label htmlFor="comment" className="block text-sm font-medium text-gray-700">
                    Comment
                </label>
                <textarea
                    id="comment"
                    name="content"
                    value={content}
                    onChange={(e) => setContent(e.target.value)}
                    placeholder="Write your comment here..."
                    required
                    className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 min-h-[100px]"
                />
            </div>
            <button
                type="submit"
                className="w-full px-4 py-2 text-white font-semibold rounded-md bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            >
                Post Comment
            </button>
        </form>
    );
}
