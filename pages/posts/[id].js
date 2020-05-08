import Layout from '../../components/layout';
import Date from '../../components/date';
import { getAllPostIds, getPostData } from '../../lib/posts';

export async function getStaticPaths() {
    // Return a list of possible value for id
    const paths = getAllPostIds();

    return {
        paths,
        fallback: false
    };
}

export async function getStaticProps({ params }) {
    // Fetch necessary data for the blog post using params.id
    const postData = await getPostData(params.id);

    return {
        props: {
            postData
        }
    };
}

function Post({ postData }) {
    return (
        <Layout title={postData.title}>
            <article>
                <div className="card mt-4 mx-auto xl:w-1/2 w-5/6">
                    <h1 className="title">{postData.title}</h1>
                    <br />
                    <div className="mb-2">
                        <Date dateString={postData.date} />
                    </div>
                    <div
                        className="post__markdown-content"
                        dangerouslySetInnerHTML={{
                            __html: postData.contentHtml
                        }}
                    />
                </div>
            </article>
        </Layout>
    );
}

export default Post;
