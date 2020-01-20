const sveltePreprocess = require("svelte-preprocess");
import { preprocessConfig } from "./rollup.config";
module.exports = {
  preprocess: preprocessConfig
};
