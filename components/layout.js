import Head from 'next/head';
import Link from 'next/link';
import Nav from './nav';

const name = 'Marco Marassi';
export const siteTitle = 'Marco Marassi - Blog';

function Layout({ children, home, title }) {
    return (
        <>
            <Head>
                <link rel="icon" href="/favicon.ico" />
                <meta
                    name="description"
                    content={home ? 'Marco Marassi - Personal Blog' : title}
                />
                <meta name="title" content={home ? siteTitle : title} />
                <meta
                    property="og:image"
                    content={`https://og-image.now.sh/${encodeURI(
                        siteTitle
                    )}.png?theme=light&md=0&fontSize=75px&images=https%3A%2F%2Fassets.zeit.co%2Fimage%2Fupload%2Ffront%2Fassets%2Fdesign%2Fnextjs-black-logo.svg`}
                />
                <meta name="og:title" content={home ? siteTitle : title} />
                <meta name="twitter:card" content="summary_large_image" />
                <link rel="apple-touch-icon" sizes="57x57" href="/favicons/apple-icon-57x57.png" />
                <link rel="apple-touch-icon" sizes="60x60" href="/favicons/apple-icon-60x60.png" />
                <link rel="apple-touch-icon" sizes="72x72" href="/favicons/apple-icon-72x72.png" />
                <link rel="apple-touch-icon" sizes="76x76" href="/favicons/apple-icon-76x76.png" />
                <link rel="apple-touch-icon" sizes="114x114" href="/favicons/apple-icon-114x114.png" />
                <link rel="apple-touch-icon" sizes="120x120" href="/favicons/apple-icon-120x120.png" />
                <link rel="apple-touch-icon" sizes="144x144" href="/favicons/apple-icon-144x144.png" />
                <link rel="apple-touch-icon" sizes="152x152" href="/favicons/apple-icon-152x152.png" />
                <link rel="apple-touch-icon" sizes="180x180" href="/favicons/apple-icon-180x180.png" />
                <link rel="icon" type="image/png" sizes="192x192"  href="/favicons/android-icon-192x192.png" />
                <link rel="icon" type="image/png" sizes="32x32" href="/favicons/favicon-32x32.png" />
                <link rel="icon" type="image/png" sizes="96x96" href="/favicons/favicon-96x96.png" />
                <link rel="icon" type="image/png" sizes="16x16" href="/favicons/favicon-16x16.png" />
                <link rel="manifest" href="/manifest.json" />
                <meta name="msapplication-TileImage" content="/favicons/ms-icon-144x144.png" />
                <meta name="msapplication-TileColor" content="#000000" />
                <meta name="theme-color" content="#000000" />
                <title>{home ? siteTitle : title}</title>
            </Head>
            <div>
                <Nav />
                <header className="header">
                    {home ? (
                        <>
                            <div>
                                <img
                                    src="/images/profile.png"
                                    alt={name}
                                    className="mx-auto w-1/6 my-4"
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
                    <footer>
                        <div className="my-4 mx-auto xl:w-1/2 w-5/6 pl-2">
                            <p className="text-gray-700">
                                &copy; {(new Date()).getFullYear()} Marco Marassi
                            </p>
                        </div>
                    </footer>
                )}
            </div>
        </>
    );
}

export default Layout;
