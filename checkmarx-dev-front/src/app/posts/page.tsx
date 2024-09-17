import Header from "@ui/Header";
import PostList from "@comp/Post/PostList";
import { fetchUser } from "app/auth/singIn";
import { fetchPosts } from "app/data/posts";
import { redirect } from "next/navigation";
import { Post, User } from "@t/index";

export default async function Posts() {
    let user: User | null = null,
        posts: Post[] | null = null;

    try {
        user = await fetchUser();
        posts = await fetchPosts();
    } catch (e) {
        console.error(e);
    }

    if (!user?.id) {
        redirect("/");
    }

    return (
        <>
            {user && <Header user={user} />}
            {user && posts && <PostList posts={posts} user={user} />};
        </>
    );
}
