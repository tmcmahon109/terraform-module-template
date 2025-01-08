#Install ASDF for package management.
if [ -f $HOME/.asdf/asdf.sh ]; then
    /bin/echo 'File exists...skipping'
else
    /usr/bin/git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.15.0
fi

# Add to zshrc if using mac
/bin/echo "Adding ASDF to .zshrc file."
/bin/echo ". $HOME/.asdf/asdf.sh" > $HOME/.zshrc

# Restart Shell
/bin/echo "Restarting shell"
exec zsh -l

# Installing plugins
 /bin/zsh install_plugins.sh
