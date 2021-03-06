# Documentation

## Using GitHub to host your blog

The service [GitHub](https://github.com) allows you to host a static website.
The site can either be delivered as markdown (.md) documents or plain HTML.
This blog generator creates plain HTML, meaning you can just upload the result
to your GitHub account and your blog will be online.

First, you'll have to [make a new GitHub repository](https://github.com/new).

The repository should be public and have your GitHub name as it's name.
For example, if your name was `hackerman`, you'd have to call the repository
`hackerman` as well. Later on, this will result in your site being published
under `https://hackerman.github.io`. If you choose any name that differs from
your account name, the address will be `https://hackerman.github.io/name`.

Next, we build the blog. As for now, you can use the `example` folder on this
repository. It contains an already buildable blog that is hostable on GitHub
as is. The only thing you might want to change, is the `config.json` at the
root of the example folder.

Next you create a new folder somewhere on your computer, but not in the
stasi-blog source folder. For example `~/home/user/blog` or
`C:/Users/user/blog`. From now on, the documentation will only be using the
path `~/home/user/...` in it's examples.

This folder now needs to be a so called "git repository".
In order to do that, you need to [download git](https://git-scm.com/downloads).

Once git has been properly installed, open a terminal and type:

```
cd YOUR_NEW_FOLDER_PATH
git clone https://github.com/hackerman/hackerman.git
```

Next you copy the example folders content to a new folder. For example
`~/home/user/blog-source`.

So the structure of that folder should now look like this:

```plain
~/home/user/blog-source
|--media
|--pages
|  |--about.html
|--articles
|  |--post-one.html
|--config.json
|--favicon.ico/png
|--README.md
```

Next you run `stasi-blog`:

```
cd ~/home/user/blog-source
stasi-blog build ./ --output="../blog"
```

This will produce the webpage on your computer. All that's left, is to move
it onto GitHub. This is done via git.

```
cd ~/home/user/blog
git add *
git commit -m "Testing my new blog"
git push
```

This procedure will be the same for every change you are doing on your blog.
I recommend to read up on git a bit, as in principle it's not too hard to
understand the basics and will probably help you in the future.

Now, the last step you have to do, is to go to your new GitHub repository and
activate the web page. Under `Settings` in the `GitHub Pages` section you have
to choose the `master` branch as the pages `source` and hit `Save`.

## Writing an article

Articles are currently written with plain HTML and require some meta
information. A document for an article should look like this:

```
{{define "title"}}Clickbait Title{{end}}
{{define "description"}}My thoughts on X.{{end}}
{{define "date"}}2020-12-10{{end}}
{{define "tags"}}categoryA,categoryB{{end}}
{{define "content"}}<p>TEXT</p>{{end}}
```

The sections `tags` and `description` are optional.

However, even though the `content` section is HTML, you don't need to write
a full web page. Instead, just write the text you'd normally want to see in
the content section of your article. While you usually start with a
heading, this can be omitted, as the heading is auto-generated by using the
`title` data.

## Writing a custom page

Writing a custom page is similar to writing an article, the only difference
is, that you only require the sections `title` and `content`.

Meaning, that this would be enough already:

```
{{define "title"}}Custom Page Name{{end}}
{{define "content"}}<p>TEXT</p>{{end}}
```

## Tables with a rowhreader

Usually a table is divided into rows and columns, where each column has a
header and each row holds data. However, there are scenarios where you want
a table to have both a column header and a row header. You can do this using
the `trh-table` style class. An example would be:

```html
<p>Opening times</p>

<table class="trh-table">
    <!--Column headers-->
    <tr>
        <!--Left empty as this cell isn't meant to be used.-->
        <th></th>
        <th>Monday</th>
        <th>Tuesday</th>
        <th>Wednesday</th>
        <th>Thursday</th>
        <th>Friday</th>
    </tr>
    <!--Beginning of data-->
    <tr>
        <!-- Rowheader cell-->
        <td>09:00 - 13:00</td>
        <!-- Data cells-->
        <td>Open</td>
        <td>Open</td>
        <td>Open</td>
        <td>Open</td>
        <td>Open</td>
    </tr>
    <tr>
        <!-- Rowheader cell-->
        <td>13:00 - 14:00</td>
        <!-- Data cells-->
        <td>Closed</td>
        <td>Closed</td>
        <td>Closed</td>
        <td>Closed</td>
        <td>Closed</td>
    </tr>
    <tr>
        <!-- Rowheader cell-->
        <td>14:00 - 18:00</td>
        <!-- Data cells-->
        <td>Open</td>
        <td>Open</td>
        <td>Open</td>
        <td>Open</td>
        <td>Closed</td>
    </tr>
</table>
```

## Best practices

### Image loading

If you want to add images to your posts, try loading them lazily, as it
allows your readers to get a readable page faster. Especially on mobile
devices, old computers or devices with a slow internet connection, this
can really help.

To lazily load an image, you need the `loading="lazy"` attribute.

An example could look like this:

```html
<img src="/images/postA/house.png" loading="lazy" alt="My new house"/>
```

### Avoid unnecessary load

Each font, script, image, video, audio file or whatever you add to your
page will cause higher load times and potentially data usage, which can
potentially even cost your user more money.

By default all fonts are default fonts and no images are used, not even
a favicon.
