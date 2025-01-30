"use client";
import { useState } from "react";
import { Input } from "@headlessui/react";
import NavBar from "../navbar/NavBar";
import FormAddnoteBtn from "../FormAddButton/AddButton";

export default function Results() {
  const [showLogin, setShowLogin] = useState<boolean>(false);
  const [formData, setFormData] = useState({ title: "", content: "" });
  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log("Form submitted with:", formData);
  };

  return (
    <>
      <NavBar setShowLogin={setShowLogin} />
      <div className="max-w-lg mx-auto p-6 sm:p-8 bg-slate-300 mt-24 flex flex-col items-center">
        <h1 className="text-center text-3xl text-slate-500 mb-6">
          Personal Note Form
        </h1>
        <form onSubmit={handleSubmit} className="w-full space-y-4">
          <div className="space-y-2">
            <label className="block text-gray-500 font-medium">Title</label>
            <Input
              className="block w-full px-3 py-2 text-gray-500 border border-gray-300 rounded-md focus:outline-none focus:ring-primary-500 focus:border-primary-500 sm:text-sm"
              placeholder="Enter Title"
              value={formData.title}
              onChange={(e) =>
                setFormData({ ...formData, title: e.target.value })
              }
            />
          </div>

          <div className="space-y-2">
            <label className="block text-gray-500 font-medium">Content</label>
            <Input
              className="block w-full px-3 py-2 text-gray-700 border border-gray-300 rounded-md focus:outline-none focus:ring-primary-500 focus:border-primary-500 sm:text-sm"
              placeholder="Enter Content"
              value={formData.content}
              onChange={(e) =>
                setFormData({ ...formData, content: e.target.value })
              }
            />
          </div>

          <div className="flex justify-center mt-6">
            <FormAddnoteBtn>Add Note</FormAddnoteBtn>
          </div>
        </form>
      </div>
    </>
  );
}
