import Header from "@ui/Header";
import PostList from "@comp/Post/PostList";
import { fetchUser } from "app/auth/signIn";
import { fetchPosts } from "app/data/posts";
import { redirect } from "next/navigation";
import { Post, User } from "@t/index";

export default async function Posts() {
    const user: User = await fetchUser();
    const posts: Post[] = await fetchPosts();

    if (!user?.id) {
        redirect("/");
    }

    return (
        <>
            <Header user={user} />
            <PostList posts={posts} user={user} />;
        </>
    );
}
