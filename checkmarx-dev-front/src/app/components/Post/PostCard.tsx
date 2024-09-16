"use client";
import PostModal from "../../ui/PostModal";
import { useRouter } from "next/navigation";
import { Pencil, Trash2 } from "lucide-react";
import { useState } from "react";
import { deletePost, updatePost } from "app/data/posts";
import { User, Post } from "@t/index";

interface PostCardProps {
    post: Post;
    user: User;
}

export default function PostCard({ post, user }: PostCardProps) {
    const [isModalOpen, setIsModalOpen] = useState(false);
    const router = useRouter();

    function handleModal() {
        setIsModalOpen((prev) => !prev);
    }

    async function handleDelete() {
        try {
            const res = await deletePost(post.id);

            if (res) {
                router.push("/posts");
            }
        } catch (e) {
            console.error(e);
        }
    }

    return (
        <article className="bg-white p-6 rounded-lg shadow-md mb-6">
            <PostModal isOpen={isModalOpen} onClose={handleModal} variant="Update" onSubmission={updatePost} postId={post.id} />
            <div className="flex justify-between items-start mb-4">
                <h1 className="text-2xl sm:text-3xl font-bold text-gray-800 break-words whitespace-normal pr-4 mb-2 sm:mb-0 w-full sm:w-auto">
                    {post.title}
                </h1>
                {user?.id === +post.author_id && (
                    <div className="flex space-x-2 mt-2 sm:mt-0 flex-shrink-0">
                        <button
                            onClick={handleModal}
                            className="p-2 rounded-full hover:bg-gray-100 transition-colors duration-200"
                        >
                            <Pencil size={24} className={isModalOpen ? "text-blue-600" : "text-gray-500"} />
                        </button>
                        <button
                            onClick={handleDelete}
                            className="p-2 rounded-full hover:bg-gray-100 transition-colors duration-200"
                        >
                            <Trash2 size={24} className="text-red-600" />
                        </button>
                    </div>
                )}
            </div>
            <p className="text-gray-600 mb-4 break-words whitespace-normal">{post.content}</p>
            <p className="text-gray-500 text-sm flex flex-wrap items-center">
                <span className="mr-2">By {user?.id === +post.author_id ? "You" : post.author_id}</span>
                <span className="mr-2">•</span>
                <span>{new Date(post.created_at).toLocaleDateString()}</span>
            </p>
        </article>
    );
}
