import Link from 'next/link';

const links = [
    { href: 'https://github.com/thtg88/blog.marco-marassi.com', label: 'Source' },
    { href: 'https://github.com/thtg88', label: 'GitHub' },
    { href: 'https://www.marco-marassi.com', label: 'Website' },
];

function Nav() {
    return (
        <nav>
            <ul className="flex justify-between items-center p-8">
                <li>
                    <Link href="/">
                        <a className="underline">Home</a>
                    </Link>
                </li>
                <li>
                    <ul className="flex justify-between items-center space-x-4">
                        {links.map(({ href, label }) => (
                            <li key={`${href}`}>
                            <a href={href} className="btn-black no-underline">
                            {label}
                            </a>
                            </li>
                        ))}
                    </ul>
                </li>
            </ul>
        </nav>
    );
}

export default Nav;
