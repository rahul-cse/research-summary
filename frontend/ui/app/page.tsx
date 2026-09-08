
"use client";

import { ChangeEvent, useState } from "react";

export default function Home() {
  const [file, setFile] = useState<File | null>(null);
  const [analysis, setAnalysis] = useState<Analysis | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  // Runs when the user selects a PDF
  const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    const selectedFile = e.target.files?.[0] ?? null;

    setFile(selectedFile);
    setAnalysis(null);
    setError("");
  };

  // Sends the PDF to the Go backend
  const handleUpload = async () => {
    if (!file) {
      setError("Please select a PDF file.");
      return;
    }

    setLoading(true);
    setError("");
    setAnalysis(null);

    // FormData is used because we are uploading a file
    const formData = new FormData();
    formData.append("file", file);

    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/upload`, {
        method: "POST",
        body: formData,
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || "Upload failed");
      }

      // Get summarized text returned by Go backend
      setAnalysis(data.text);
      console.log(analysis)

    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : "Something went wrong."
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <main className="min-h-screen bg-gray-100 p-8">
      <div className="mx-auto max-w-5xl rounded-xl bg-white p-8 shadow">

        {/* Page heading */}
        <h1 className="text-3xl font-bold">
          Research Paper Summarizer
        </h1>

        <p className="mt-2 text-gray-600">
          Upload a research paper to extract and produce a summarization.
        </p>

        {/* Upload area */}
        <div className="mt-8 rounded-lg border-2 border-dashed p-8">
          <div className="flex items-center gap-4">
            <label
              htmlFor="pdf-upload"
              className="inline-block cursor-pointer rounded-lg border px-5 py-2">
                Choose PDF
            </label>
          <input
            id="pdf-upload"
            type="file"
            accept="application/pdf,.pdf"
            onChange={handleFileChange}
            className="hidden"
          />

          {/* Selected file name */}
          {file && (
            <p className="mt-3 text-sm text-gray-600">
              Selected file: {file.name}
            </p>
          )}

          {/* Upload button */}
            <div className="flex justify-center">
                <button
                 onClick={handleUpload}
                 disabled={loading || !file}
                 className="rounded-lg bg-black px-5 py-2 text-white disabled:cursor-not-allowed disabled:opacity-50">
                 {loading ? "Processing..." : "Upload & Process"}
              </button>
            </div>
            

          {/* Error message */}
          {error && (
            <p className="mt-4 text-red-600">
              {error}
            </p>
          )}
        </div>
        </div>
        {/* Extracted text */}
        {analysis && (
          <section className="mt-8">

            <h2 className="mb-3 text-xl font-semibold">
              Paper Summary
            </h2>

            <div className="max-h-[600px] overflow-y-auto whitespace-pre-wrap rounded-lg border bg-gray-50 p-5 text-sm leading-7">
                <div className="mb-6">
                  <h3 className="mb-2 font-semibold">
                    Research Domain
                  </h3>

                  <div className="flex flex-wrap gap-2">
                    {
                      analysis.domain.map((item)=> <span key={item} className="rounded-full bg-white px-3 py-1 text-sm border">
                      {item} </span>)
                    }
                    
                  </div>
                  <h3 className="mb-2 font-semibold">
                    Methods
                  </h3>
                  <div className="flex flex-wrap gap-2">
                      {analysis.methods.map((item) => (
                      <span key={item} className="rounded-full bg-white px-3 py-1 text-sm border">
                      {item} </span>
                      ))}
                  </div>

                  <h3 className="mb-2 font-semibold">
                    Study Application
                  </h3>
                  <div className="flex flex-wrap gap-2">
                      {analysis.study_application.map((item) => (
                      <span key={item} className="rounded-full bg-white px-3 py-1 text-sm border">
                      {item} </span>
                      ))}
                  </div>

                  <h3 className="mb-2 font-semibold">
                    Summary
                  </h3>
                  <div>
                    {analysis.summary}
                  </div>
                </div>
            
            </div>

          </section>
        )}
      </div>
    </main>
  );
}