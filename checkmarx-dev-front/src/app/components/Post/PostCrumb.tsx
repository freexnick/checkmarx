"use client";
import Link from "next/link";
import { Trash2 } from "lucide-react";
import { deletePost } from "app/data/posts";
import { useRouter } from "next/navigation";
import { User, Post } from "@t/index";

interface PostCrumbProps {
    user: User;
    post: Post;
}

export default function PostCrumb({ user, post }: PostCrumbProps) {
    const router = useRouter();

    async function handleDelete() {
        try {
            const res = await deletePost(post.id);
            if (res) {
                router.refresh();
            }
        } catch (e) {
            console.error(e);
        }
    }

    return (
        <div className="flex">
            <div className="p-3 flex-grow">
                <div className="flex items-center text-xs text-gray-500">
                    <div className="flex justify-between items-center w-full">
                        <span>
                            By {user?.id === +post.author_id ? "You" : post.author_id} •&nbsp;
                            {new Date(post.created_at).toLocaleDateString()}
                        </span>
                        {user?.id === +post.author_id && (
                            <button
                                onClick={handleDelete}
                                className="p-2 rounded-full hover:bg-gray-100 transition-colors duration-200"
                            >
                                <Trash2 color="red" size={16} />
                            </button>
                        )}
                    </div>
                </div>
                <Link href={`/posts/${post.id}`} className="text-lg font-semibold mb-2">
                    {post.title}
                </Link>
                <p className="text-sm text-gray-700 mb-2">{post.content}</p>
            </div>
        </div>
    );
}
