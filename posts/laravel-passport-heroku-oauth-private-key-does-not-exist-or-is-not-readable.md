---
title: 'Laravel Passport on Heroku - OAuth Private Key Does Not Exist or Is Not Readable'
date: '2020-05-10'
---

Recently I was deploying a [Laravel 7](https://laravel.com/) REST API, to a dyno on Heroku.
This particular project uses [Passport](https://laravel.com/docs/7.x/passport) for authentication, so after deploying the app, I generated the OAuth keys using the `php artisan passport:keys` command.
I then tried to log in, feeling confident this should be a fairly easy, and successful setup.
Unfortunately I hit the following error:

```
Key path "file:///app/storage/oauth-private.key" does not exist or is not readable
```

I was pretty sure the `passport:keys` command succeeded, and a quick `heroku run bash` and a `ls` in the project's `storage` folder, confirmed my theory, and all environment variables (like `APP_URL`) were set correctly, and other parts of the API that didn't require authentication were working fine.
Why wasn't the login functionality working then?

The reason why this was not working is due to Heroku's ephemeral file system, where files generated from the application apart from the deployed files are not persisted and, even if they are, they get removed within at most 24 hours, as this [Heroku Help article](https://help.heroku.com/K1PPS2WM/why-are-my-file-uploads-missing-deleted) states:

> The Heroku filesystem is ephemeral - that means that any changes to the filesystem whilst the dyno is running only last until that dyno is shut down or restarted. Each dyno boots with a clean copy of the filesystem from the most recent deploy. This is similar to how many container based systems, such as Docker, operate.
> In addition, under normal operations dynos will restart every day in a process known as "Cycling".
> These two facts mean that the filesystem on Heroku is not suitable for persistent storage of data. [...]

How can you solve this issue then? Luckily Passport has the ability to specify public and private keys as environment variables.

From your local machine run `php artisan vendor:publish --tag=passport-config`.
Laravel will publish the Passport config files within your `config` folder.
Make sure you deploy these changes to Heroku.

Next, go to your [Heroku Dashboard](https://dashboard.heroku.com/apps), and put your OAuth keys as config vars in the App Settings page.
You should be able to see what these are named from the Passport config file, these are `PASSPORT_PRIVATE_KEY` and `PASSPORT_PUBLIC_KEY`.

Finally, try to login, you should now be able to login to your REST API using Laravel Passport on Heroku! :-)

For more info about deploying projects that use Passport, see the [official docs](https://laravel.com/docs/7.x/passport#deploying-passport).
