import React, { forwardRef } from "react";

export const Button = forwardRef<
    HTMLButtonElement,
    React.ButtonHTMLAttributes<HTMLButtonElement> & { variant?: "default" | "outline"; size?: "default" | "sm" }
>(({ className, variant = "default", size = "default", ...props }, ref) => {
    const baseStyles = "font-medium rounded focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500";
    const variantStyles = {
        default: "bg-blue-500 text-white hover:bg-blue-600",
        outline: "border border-gray-300 text-gray-700 hover:bg-gray-50",
    };
    const sizeStyles = {
        default: "px-4 py-2",
        sm: "px-3 py-1 text-sm",
    };

    return <button className={`${baseStyles} ${variantStyles[variant]} ${sizeStyles[size]} ${className}`} ref={ref} {...props} />;
});
Button.displayName = "Button";
