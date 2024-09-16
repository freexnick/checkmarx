"use client";
import { fetchUser } from "app/auth/signIn";
import { useRouter } from "next/navigation";
import { MutableRefObject, useRef, createContext, useContext, useLayoutEffect } from "react";
import { User } from "@t/index";

type userRefType = User | null;

type UserContextType = {
    userRef: MutableRefObject<userRefType>;
};

const UserContext = createContext<UserContextType>({
    userRef: { current: null },
});

function UserProvider({ children }: { children: React.ReactNode }) {
    const router = useRouter();
    const userRef = useRef<userRefType>(null);

    async function getUser() {
        try {
            const user = (await fetchUser()) as User;
            if (user?.id) {
                userRef.current = {
                    id: user.id,
                    email: user.email,
                };
                return;
            }
            router.replace("/");
        } catch (e) {
            console.error(e);
            router.replace("/");
        }
    }

    useLayoutEffect(() => {
        if (!userRef.current) {
            getUser();
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    return <UserContext.Provider value={{ userRef }}>{children}</UserContext.Provider>;
}

function useUser(): UserContextType {
    return useContext(UserContext);
}

export { UserProvider, useUser };
