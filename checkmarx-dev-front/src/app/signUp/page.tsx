import { API_URI } from "@conf/config";
import SignUpForm, { SubmitResult } from "@comp/Auth/SignUpForm";

async function handleSignUp(data: Record<string, FormDataEntryValue>): Promise<Awaited<SubmitResult>> {
    "use server";
    try {
        const response = await fetch(`${API_URI}/auth/signup`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),

            credentials: "include",
        });

        return response.json();
    } catch (err) {
        console.error(err);
    }
}

export default function SignUp() {
    return <SignUpForm handleSignUp={handleSignUp} />;
}
