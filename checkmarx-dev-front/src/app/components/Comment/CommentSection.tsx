"use client";
import React, { useState } from "react";
import CommentCard from "./CommentCard";
import CommentForm from "./CommentForm";
import { Comment, User, CommentState } from "@t/index";

interface CommentSectionProprs {
    comments: Comment[];
    postID: number;
    user: User;
}

export default function CommentSection({ comments, postID, user }: CommentSectionProprs) {
    const [commentState, setCommentState] = useState<CommentState | null>(null);

    function stopEditing() {
        setCommentState(null);
    }

    return (
        <>
            {comments?.length ? (
                comments.map((comment: Comment) => (
                    <React.Fragment key={comment.id}>
                        <CommentCard
                            comment={comment}
                            setCommentState={setCommentState}
                            commentState={commentState}
                            user={user}
                        />
                    </React.Fragment>
                ))
            ) : (
                <p className="text-gray-500">No comments yet.</p>
            )}
            <section>
                <CommentForm postID={postID} commentState={commentState} stopEditing={stopEditing} />
            </section>
        </>
    );
}
