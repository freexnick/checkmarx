import Post from "@comp/Post/Post";
import { Post as PostType, User } from "@t/index";

interface PostListProps {
    posts: PostType[];
    user: User;
}

export default function PostList({ posts, user }: PostListProps) {
    return (
        <>
            <div className="max-w-3xl mx-auto p-6 bg-gray-50 text-gray-900">
                <section className="space-y-4">
                    {posts ? (
                        posts?.map((post: PostType) => <Post key={post.id} post={post} user={user} />)
                    ) : (
                        <div>No Posts Yet</div>
                    )}
                </section>
            </div>
        </>
    );
}
