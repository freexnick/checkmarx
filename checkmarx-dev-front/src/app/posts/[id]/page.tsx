import Link from "next/link";
import Header from "@ui/Header";
import PostCard from "@comp/Post/PostCard";
import Spinner from "@ui/Spinner";
import CommentSection from "@comp/Comment/CommentSection";
import { fetchPost } from "app/data/posts";
import { fetchUser } from "app/auth/singIn";
import { redirect } from "next/navigation";
import { Post, Comment, User } from "@t/index";

export default async function PostPage({ params }: { params: { id: string } }) {
    let user: User | null = null;
    let post: Post | null = null;
    let comments: Comment[] = [];

    try {
        user = await fetchUser();
        const result = await fetchPost(params.id);
        if (result) {
            post = result.post;
            comments = result.comments;
        }
    } catch (e) {
        console.error(e);
    }

    if (!user?.id) {
        redirect("/");
    }

    return (
        <>
            <Header user={user} />
            <div className="max-w-3xl mx-auto p-4 sm:p-6 bg-gray-50 text-gray-900">
                {post ? (
                    <>
                        <header className="mb-4">
                            <Link href="/posts" className="text-blue-500 hover:underline">
                                &larr; Back to Posts
                            </Link>
                        </header>
                        <div className="break-words whitespace-normal">
                            <PostCard post={post} user={user} />
                        </div>
                        <section className="mb-6">
                            <h2 className="text-2xl font-semibold mb-4 text-gray-800">Comments</h2>
                            <ul className="space-y-4">
                                <CommentSection comments={comments} postID={post.id} user={user} />
                            </ul>
                        </section>
                    </>
                ) : (
                    <Spinner />
                )}
            </div>
        </>
    );
}
