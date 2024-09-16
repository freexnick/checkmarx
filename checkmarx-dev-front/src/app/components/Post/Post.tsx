import React from "react";
import PostCrumb from "./PostCrumb";
import { Post as PostType, User } from "@t/index";

interface PostProps {
    post: PostType;
    user: User;
}

export default function Post({ post, user }: PostProps) {
    return (
        <article className="bg-white rounded-md shadow-sm overflow-hidden text-pretty text-wrap:wrap break-normal">
            <PostCrumb post={post} user={user} />
        </article>
    );
}
