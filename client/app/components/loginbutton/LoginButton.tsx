import { Button } from "@headlessui/react";
interface LoginButton {
  children: React.ReactNode;

}
export default function LoginButton({ children }: LoginButton) {
  return (
    <Button
      className=" rounded-full px-2 py-1 md:px-8 md:py-2 bg-transparent text-accent border border-accent hover:bg-accent hover:text-white cursor-pointer"
    >
      {children}
    </Button>
  );
}
