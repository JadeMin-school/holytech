import type { Metadata } from "next";



export const metadata: Metadata = {
	title: "Holytech",
	description: "폴리텍 정보 앱",
};

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	return (
		<html lang="ko">
			<body>{children}</body>
		</html>
	);
};