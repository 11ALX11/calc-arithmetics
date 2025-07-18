#! /bin/sh

initialize_locale_directories() {
    # Iterate over each provided language code
    for lang_code in "$@"; do
        # Define locale directory path
        locale_dir="locales/$lang_code/LC_MESSAGES"

        # Create the locale directory if it doesn't exist
        mkdir -p "$locale_dir"

        # Initialize the default.po file using msginit
        msginit --no-translator -o "$locale_dir/default.po" \
            -l "$lang_code" -i "locales/default.po"

        # Print status message
        echo "Initialized locale directory"
        echo "for $lang_code"
    done
}

# Check if any language codes are provided as arguments
if [ $# -eq 0 ]; then
    echo "Error: No language codes provided."
    echo "Please provide one or more language codes as arguments."
    exit 1
fi

initialize_locale_directories "$@"