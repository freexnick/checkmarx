"use client";
import Link from "next/link";
import { AnimatePresence, motion } from "framer-motion";
import { useState } from "react";
import { EyeIcon, EyeOff, Mail } from "lucide-react";
import { useRouter } from "next/navigation";
import { useUser } from "app/context/User";
import { signUser } from "app/auth/singIn";

export default function SignInForm() {
    const [isPasswordVisible, setIsPasswordVisible] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const { userRef } = useUser();
    const router = useRouter();

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();
        setErrorMessage("");
        const formData = new FormData(event.currentTarget);
        const data = Object.fromEntries(formData.entries());

        try {
            const response = await signUser(data);
            if (response?.email) {
                userRef.current = { id: +response.user_id, email: response.email };
                router.push("/posts");
            }
            setErrorMessage(response?.message);
        } catch (e) {
            console.error(e);
        }
    }

    const togglePasswordVisibility = () => {
        setIsPasswordVisible((prev) => !prev);
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
            <AnimatePresence>
                <motion.div
                    initial={{ opacity: 0, y: -50 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ duration: 0.5 }}
                    className="bg-white h-sign p-8 rounded-lg shadow-md w-96"
                >
                    <motion.h1
                        initial={{ opacity: 0, y: -50 }}
                        animate={{ opacity: 1, y: 0 }}
                        transition={{ duration: 0.5 }}
                        className="text-2xl font-bold mb-10 text-center text-gray-800"
                    >
                        Sign In
                    </motion.h1>
                    <form className="rounded px-8 py-8 mb-4 group" onSubmit={handleSubmit}>
                        <motion.div
                            initial={{ opacity: 0, y: -50 }}
                            animate={{ opacity: 1, x: 0 }}
                            transition={{ delay: 0.3, duration: 0.5 }}
                            className="mb-2 h-24"
                        >
                            <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="email">
                                Email
                            </label>
                            <input
                                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline focus:ring-2 focus:ring-blue-500 invalid:[&:not(:placeholder-shown):not(:focus)]:outline-border-red-500 peer"
                                id="email"
                                name="email"
                                type="email"
                                placeholder="Email"
                                required
                                minLength={6}
                                maxLength={255}
                            />
                            <button
                                type="button"
                                onClick={togglePasswordVisibility}
                                className="h-5 w-5 absolute inset-y-0 right-3 top-9 flex items-center justify-center text-gray-500 hover:text-blue-500 focus:outline-none"
                            >
                                <Mail />
                            </button>
                            <motion.p
                                initial={{ opacity: 0 }}
                                animate={{ opacity: 1 }}
                                className="mt-2 hidden text-sm text-red-500 peer-[&:not(:placeholder-shown):not(:focus):invalid]:block"
                            >
                                Please enter a valid email address
                            </motion.p>
                        </motion.div>
                        <motion.div
                            initial={{ opacity: 0, y: -50 }}
                            animate={{ opacity: 1, x: 0 }}
                            transition={{ delay: 0.4, duration: 0.5 }}
                            className="mb-6 h-24"
                        >
                            <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="password">
                                Password
                            </label>
                            <input
                                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline focus:ring-2 focus:ring-blue-500 invalid:[&:not(:placeholder-shown):not(:focus)]:outline-border-red-500 peer"
                                id="password"
                                name="password"
                                type={isPasswordVisible ? "text" : "password"}
                                placeholder="********"
                                minLength={8}
                                maxLength={255}
                                required
                            />
                            <button
                                type="button"
                                onClick={togglePasswordVisibility}
                                className="h-5 w-5 absolute inset-y-0 right-3 top-9 flex items-center justify-center text-gray-500 hover:text-blue-500 focus:outline-none"
                            >
                                {isPasswordVisible ? <EyeOff /> : <EyeIcon />}
                            </button>
                            <motion.p
                                initial={{ opacity: 0 }}
                                animate={{ opacity: 1 }}
                                className="my-2 hidden text-sm text-red-500 peer-[&:not(:placeholder-shown):not(:focus):invalid]:block"
                            >
                                Password must be 8 characters
                            </motion.p>
                            {errorMessage && <motion.p className=" text-sm text-red-500">{errorMessage}</motion.p>}
                        </motion.div>
                        <div className="mt-2.5">
                            <motion.div
                                initial={{ opacity: 0, y: -50 }}
                                animate={{ opacity: 1, x: 0 }}
                                transition={{ delay: 0.5, duration: 0.5 }}
                                className="flex items-center justify-center w-full"
                            >
                                <button
                                    className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 my-2.5 rounded focus:outline-none focus:shadow-outline group-invalid:pointer-events-none group-invalid:opacity-30 w-full"
                                    type="submit"
                                >
                                    Sign In
                                </button>
                            </motion.div>
                            <motion.div
                                initial={{ opacity: 0, y: -50 }}
                                animate={{ opacity: 1, x: 0 }}
                                transition={{ delay: 0.6, duration: 0.5 }}
                                className="mt-4 text-center text-sm text-gray-600"
                            >
                                Don&apos;t have an account?
                                <Link href="/signUp" className="text-blue-500 hover:underline ml-1">
                                    Sign Up
                                </Link>
                            </motion.div>
                        </div>
                    </form>
                </motion.div>
            </AnimatePresence>
        </div>
    );
}
