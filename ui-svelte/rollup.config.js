import svelte from "rollup-plugin-svelte";
import resolve from "@rollup/plugin-node-resolve";
import commonjs from "@rollup/plugin-commonjs";
import livereload from "rollup-plugin-livereload";
import { terser } from "rollup-plugin-terser";
import sveltePreprocess from "svelte-preprocess";
const production = !process.env.ROLLUP_WATCH;

export const preprocess = sveltePreprocess({
  scss: {
    includePaths: ["src"],
    data: `@import 'scss/_variables.scss';`
  },
  postcss: {
    plugins: [require("autoprefixer")]
  },
  typescript: {
    /**
     * Optionally specify the directory to load the tsconfig from
     */
    tsconfigDirectory: "./configs",

    /**
     * Optionally specify the full path to the tsconfig
     */
    tsconfigFile: "./tsconfig.app.json",

    /**
     * Optionally specify compiler options.
     * These will be merged with options from the tsconfig if found.
     */
    compilerOptions: {
      module: "es2015"
    },

    /**
     * Type checking can be skipped by setting 'transpileOnly: true'.
     * This speeds up your build process.
     */
    transpileOnly: true
  }
});

export default {
  input: "src/main.js",
  output: {
    sourcemap: true,
    format: "iife",
    name: "app",
    file: "public/build/bundle.js"
  },
  plugins: [
    svelte({
      // enable run-time checks when not in production
      dev: !production,
      // we'll extract any component CSS out into
      // a separate file — better for performance
      css: css => {
        css.write("public/build/bundle.css");
      },
      preprocess
    }),

    // If you have external dependencies installed from
    // npm, you'll most likely need these plugins. In
    // some cases you'll need additional configuration —
    // consult the documentation for details:
    // https://github.com/rollup/plugins/tree/master/packages/commonjs
    resolve({
      browser: true,
      dedupe: importee =>
        importee === "svelte" || importee.startsWith("svelte/")
    }),
    commonjs(),

    // In dev mode, call `npm run start` once
    // the bundle has been generated
    !production && serve(),

    // Watch the `public` directory and refresh the
    // browser on changes when not in production
    !production && livereload("public"),

    // If we're building for production (npm run build
    // instead of npm run dev), minify
    production && terser()
  ],
  watch: {
    clearScreen: false
  }
};

function serve() {
  let started = false;

  return {
    writeBundle() {
      if (!started) {
        started = true;

        require("child_process").spawn(
          "npm",
          ["run", "start", "--", "--single", "--dev"],
          {
            stdio: ["ignore", "inherit", "inherit"],
            shell: true
          }
        );
      }
    }
  };
}
