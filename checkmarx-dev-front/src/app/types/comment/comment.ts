type Comment = {
    id: number;
    content: string;
    author_id: number;
    created_at: string;
};

type CommentState = {
    commentID: number;
    commentMessage: string;
};
export type { Comment, CommentState };
