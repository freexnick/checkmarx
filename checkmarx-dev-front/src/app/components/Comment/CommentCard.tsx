"use client";
import { useRouter } from "next/navigation";
import { API_URI } from "@conf/config";
import { Pencil, Trash2 } from "lucide-react";
import { User, CommentState, Comment } from "@t/index";

interface CommentCardProps {
    comment: Comment;
    commentState: CommentState | null;
    setCommentState: React.Dispatch<React.SetStateAction<CommentState | null>>;
    user: User;
}

export default function CommentCard({ comment, setCommentState, commentState, user }: CommentCardProps) {
    const router = useRouter();
    const logged_user = user?.id;
    const author = logged_user || user?.id;
    const owner = logged_user === comment.author_id;

    function handleEdit(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
        e.stopPropagation();
        const current = commentState?.commentID == comment.id;
        setCommentState(current ? null : { commentID: comment.id, commentMessage: comment.content });
    }

    async function handleDelete() {
        try {
            const response = await fetch(`${API_URI}/comments/${comment.id}`, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
            });

            if (response.ok) {
                router.refresh();
            }
        } catch (e) {
            console.error(e);
        }
    }

    return (
        <li key={comment.id} className="bg-white p-4 rounded-lg shadow-md">
            <div className="flex justify-between">
                <span className="font-semibold text-gray-700">{owner ? "You" : author}</span>
                {owner && (
                    <div>
                        <button
                            onClick={handleEdit}
                            className="p-2 rounded-full hover:bg-gray-100 transition-colors duration-200"
                        >
                            <Pencil
                                size={18}
                                className={commentState?.commentID === comment.id ? "text-blue-600" : "text-gray-500"}
                            />
                        </button>
                        <button
                            onClick={handleDelete}
                            className="p-2 rounded-full hover:bg-gray-100 transition-colors duration-200"
                        >
                            <Trash2 size={20} color="red" />
                        </button>
                    </div>
                )}
            </div>
            <p className="mt-1 text-gray-800">{comment.content}</p>
            <p className="text-sm text-gray-500 mt-2">{new Date(comment.created_at).toLocaleDateString()}</p>
        </li>
    );
}
