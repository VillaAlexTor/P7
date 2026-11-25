import type { Metadata } from "next";
import "@fontsource-variable/manrope";
import "@fontsource-variable/source-serif-4";
import "./globals.css";

export const metadata: Metadata = {
  title: "p7 — Privacidad, descifrada",
  description: "Evaluaciones de privacidad personalizadas, explicables y respaldadas por evidencias.",
};

export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="es">
      <body>{children}</body>
    </html>
  );
}

