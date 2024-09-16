import { ReactNode } from "react";
import { UserProvider } from "./User";

export default function Providers({ children }: { children: ReactNode }) {
    return <UserProvider>{children}</UserProvider>;
}
