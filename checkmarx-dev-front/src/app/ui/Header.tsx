"use client";
import PostModal from "@ui/PostModal";
import Link from "next/link";
import { useState } from "react";
import { Plus } from "lucide-react";
import { Button } from "@ui/Button";
import { createPost } from "app/data/posts";
import { User } from "@t/index";
import { signOut } from "app/auth/singOut";
import { useRouter } from "next/navigation";

interface HeaderProps {
    user: User;
}

export default function Header({ user }: HeaderProps) {
    const [isModalOpen, setIsModalOpen] = useState(false);
    const router = useRouter();

    function handleModal() {
        setIsModalOpen((prev) => !prev);
    }

    function handleSignOut() {
        signOut();
        router.push("/");
    }

    return (
        <div className="bg-white border-b border-gray-200">
            <div className="max-w-4xl mx-auto px-4 py-2 flex items-center justify-between">
                <Link href={"/posts"} className="text-2xl font-bold text-gray-800">
                    Fresh Posts
                </Link>
                <div className="flex gap-2">
                    <Button onClick={() => setIsModalOpen(true)} variant="default" className="flex items-center space-x-2">
                        <Plus size={20} />
                        <span>Create Post</span>
                    </Button>
                    <span className="mr-2 py-2 font-semibold text-blue-500">Logged as {user?.email}</span>
                    <Button type="button" variant="outline" onClick={handleSignOut}>
                        Logout
                    </Button>
                </div>
            </div>
            <PostModal isOpen={isModalOpen} onClose={handleModal} variant="Create" onSubmission={createPost} />
        </div>
    );
}
