import { Button } from "@headlessui/react";

interface AddnoteBtnProps {
  children: React.ReactNode;
}

export default function FormAddnoteBtn({ children }: AddnoteBtnProps) {
  return (
    <Button className="flex text-lg rounded-full px-6 py-3 md:px-6 md:py-3 bg-transparent text-accent border border-accent hover:bg-accent hover:text-white cursor-pointer mt-10 items-center justify-center">
      {children}
    </Button>
  );
}
