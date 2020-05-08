import Head from 'next/head';
import Link from 'next/link';
import Layout, { siteTitle } from '../components/layout';
import { getSortedPostsData } from '../lib/posts';
import Date from '../components/date';

export async function getStaticProps() {
    // // Instead of the file system,
    // // fetch post data from an external API endpoint
    // const res = await fetch('..')
    // return res.json()
    const allPostsData = getSortedPostsData();

    return {
        props: {
            allPostsData
        }
    };
}

function Home({ allPostsData }) {
    return (
        <Layout home>
            <Head><title>{siteTitle}</title></Head>
            <section>
                {
                    allPostsData.length === 0 &&
                    <h2 className="subtitle text-center mt-4">Coming soon!</h2>
                }
                {allPostsData.map(({ id, date, title }) => (
                    <Link key={id} href="/posts/[id]" as={`/posts/${id}`}>
                        <a>
                            <div className="card mt-6 mx-auto sm:w-1/2 w-5/6">
                                <strong>{title}</strong>
                                <br />
                                <Date dateString={date} />
                            </div>
                        </a>
                    </Link>
                ))}
            </section>
        </Layout>
    );
}

export default Home;
