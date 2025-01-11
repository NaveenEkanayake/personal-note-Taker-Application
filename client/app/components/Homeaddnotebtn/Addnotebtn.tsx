import { Button } from "@headlessui/react";

interface AddnoteBtnProps {
  children: React.ReactNode;
}

export default function AddnoteBtn({ children }: AddnoteBtnProps) {
  return (
    <Button className="text-lg rounded-full px-4 py-2 md:px-10 md:py-4 bg-transparent text-accent border border-accent hover:bg-accent hover:text-white cursor-pointer">
      {children}
    </Button>
  );
}
