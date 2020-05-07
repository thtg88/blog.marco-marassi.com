import Head from 'next/head';
import Link from 'next/link';

const name = 'Marco Marassi';
export const siteTitle = 'Blog - Marco Marassi';

function Layout({ children, home }) {
    return (
        <>
            <Head>
                <link rel="icon" href="/favicon.ico" />
                <meta
                    name="description"
                    content="Learn how to build a personal website using Next.js"
                />
                <meta
                    property="og:image"
                    content={`https://og-image.now.sh/${encodeURI(
                        siteTitle
                    )}.png?theme=light&md=0&fontSize=75px&images=https%3A%2F%2Fassets.zeit.co%2Fimage%2Fupload%2Ffront%2Fassets%2Fdesign%2Fnextjs-black-logo.svg`}
                />
                <meta name="og:title" content={siteTitle} />
                <meta name="twitter:card" content="summary_large_image" />
            </Head>
            <div>
                <header className="header">
                    {home ? (
                        <>
                            <div>
                                <img
                                    src="/images/profile.png"
                                    alt={name}
                                    className="mx-auto w-1/6"
                                />
                            </div>
                            <h1 className="title">{siteTitle}</h1>
                        </>
                    ) : (
                        <Link href="/">
                            <a>
                                <img
                                    src="/images/profile.png"
                                    alt={name}
                                    className="mx-auto w-1/6"
                                />
                                <h2 className="subtitle">{siteTitle}</h2>
                            </a>
                        </Link>
                    )}
                </header>
                <main>{children}</main>
                {!home && (
                    <div className="my-4 mx-auto xl:w-1/2 w-5/6 pl-2">
                        <Link href="/">
                            <a className="text-blue-500">&larr; Back to home</a>
                        </Link>
                    </div>
                )}
            </div>
        </>
    );
}

export default Layout;
