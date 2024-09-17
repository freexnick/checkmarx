import { Button } from "@ui/Button";
import { useUser } from "app/context/User";
import { motion, AnimatePresence } from "framer-motion";
import { useRouter } from "next/navigation";

interface PostModalProps {
    isOpen: boolean;
    onClose: () => void;
    variant: "Create" | "Update";
    onSubmission: (data: Record<string, FormDataEntryValue | number>) => Promise<string | undefined>;
    postId?: number;
}

export default function PostModal({ isOpen, onClose, variant, onSubmission, postId }: PostModalProps) {
    const { userRef } = useUser();
    const router = useRouter();

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();
        const formData = new FormData(event.currentTarget);
        const user_id = userRef?.current?.id as number;
        const data = Object.fromEntries(formData.entries()) as Record<string, FormDataEntryValue | number>;
        data.author_id = user_id;
        if (postId) data.id = postId;

        try {
            const res = await onSubmission(data);

            if (res) {
                onClose();
                router.refresh();
            }
        } catch (e) {
            console.error(e);
        }
    }

    return (
        <AnimatePresence>
            {isOpen && (
                <motion.div
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    exit={{ opacity: 0 }}
                    className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50"
                >
                    <motion.div
                        initial={{ scale: 0.9, opacity: 0 }}
                        animate={{ scale: 1, opacity: 1 }}
                        exit={{ scale: 0.9, opacity: 0 }}
                        className="bg-white rounded-lg p-6 w-full max-w-md"
                    >
                        <h2 className="text-xl font-semibold mb-4 text-blue-500">{variant} Post</h2>
                        <form onSubmit={handleSubmit} className="space-y-4">
                            <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="post_title"></label>
                            <input
                                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline focus:ring-2 focus:ring-blue-500 invalid:[&:not(:placeholder-shown):not(:focus)]:outline-border-red-500 peer"
                                id="post_title"
                                name="title"
                                type="text"
                                placeholder="Post Title"
                                required
                                minLength={6}
                                maxLength={120}
                            />
                            <label
                                htmlFor="post_message"
                                className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                            >
                                Your message
                            </label>
                            <textarea
                                id="post_message"
                                name="content"
                                rows={4}
                                maxLength={255}
                                className="block p-2.5 w-full text-sm text-gray-900 bg-white rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500"
                                placeholder="Write your thoughts here..."
                                required
                            />
                            <div className="flex justify-end space-x-2">
                                <Button type="button" variant="outline" onClick={onClose}>
                                    Cancel
                                </Button>
                                <Button type="submit">{variant} Post</Button>
                            </div>
                        </form>
                    </motion.div>
                </motion.div>
            )}
        </AnimatePresence>
    );
}
